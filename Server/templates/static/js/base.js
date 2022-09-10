var HTTPCode = {
    Success: 200,
    BadRequest: 400,
    NotFound: 404,
    Conflict: 409,
    InternalServerError: 500,
}

var Elements = {
    LoginEmail: "loginEmail",
    LoginPassword: "loginPassword",
    LoginModalBody: "loginModalBody",
    RegisterFirstName: "registerFirstName",
    RegisterLastName: "registerLastName",
    RegisterEmail: "registerEmail",
    RegisterPassword: "registerPassword",
    RegisterPasswordConfirm: "registerPasswordConfirm",
    RegisterModalBody: "registerModalBody"
}

let x = document.getElementById("loginEmail")
x.addEventListener("keydown", function(e){
    if (e.code == "Enter") {
        login()
    }
})

let y = document.getElementById("loginPassword")
y.addEventListener("keydown", function(e){
    if (e.code == "Enter") {
        login()
    }
})

function addAlert(alertLevel, alertMessage, target) {
    let alert = 
    `
    <div class="container">
      <div class="alert alert-${alertLevel} alert-dismissible fade show mx-auto" role="alert">
        ${alertMessage}
        <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>
      </div>
    </div>
    `
    let element = document.getElementById(target)
    element.insertAdjacentHTML("afterbegin", alert)
}

async function register() {
    let errorMessage = checkRegistration()
    if (errorMessage != "") {
        addAlert("warning", errorMessage, Elements.RegisterModalBody)
        return
    }
}

function checkRegistration() {
    let listStart = `<p class="mb-0">`, listEnd = `</p>`
    let returnString = ""
    if (validName(Elements.RegisterFirstName)) {
        returnString += listStart+"First Name is empty."+listEnd
    }
    if (validName(Elements.RegisterLastName)) {
        returnString += listStart+"Last Name is empty."+listEnd
    }
    if (!validEmail(Elements.RegisterEmail)) {
        returnString += listStart+"Email format is invalid."+listEnd
    }
    if (document.getElementById(Elements.RegisterPassword).value != document.getElementById(Elements.RegisterPasswordConfirm).value) {
        returnString += listStart+"Passwords do not match."+listEnd
    }
    if (!validPassword(Elements.RegisterPassword)) {
        returnString += listStart+"Password is not an appropriate length."+listEnd
    }
    return returnString
}

async function login() {
    if (!checkLogin(Elements.LoginEmail, Elements.LoginPassword)) {
        addAlert("warning", "The provided login information are invalid.", Elements.LoginModalBody)
        console.log("validation: client")
        return
    }

    let response = await fetch("http://localhost:8888/login", 
    {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            UserID: document.getElementById(Elements.LoginEmail).value,
            Password: document.getElementById(Elements.LoginPassword).value
        })
    })

    switch (response.status) {
        case HTTPCode.Success:
            addAlert("success", "Login was successful.", Elements.LoginModalBody)
            await sleep(2000)
            location.reload()
            break
        case HTTPCode.BadRequest:
            addAlert("warning", "The provided login information are invalid. (S)", Elements.LoginModalBody)
            console.log("validation: server")
            break
        case HTTPCode.NotFound:
            addAlert("warning", "No user was found with provided credentials.", Elements.LoginModalBody)
            break
        case HTTPCode.InternalServerError:
            addAlert("danger", "Internal server error.", Elements.LoginModalBody)
            break
        default:
            addAlert("danger", "Unknown error occured.", Elements.LoginModalBody)
            console.log("Login attempt status code: " + response.status)
            break
    }
}

// currently only updates login button based on focus changing from one of the text input fields
// possibly find another solution, or instead remove disable part and instead show a warning if login is tried without acceptable input
function checkLogin() {
    if (validEmail(Elements.LoginEmail) && validPassword(Elements.LoginPassword)) {
        return true
    } else {
        return false
    }
}

// only checks for @ and non empty
function validEmail(targetID) {
    let email = document.getElementById(targetID).value
    if (email == "" || !email.includes("@")) {
        return false
    }
    return true
}

// only checkcs length
function validPassword(targetID) {
    let password = document.getElementById(targetID).value
    if (password.length >= 8 && password.length <= 32) {
        return true
    } else {
        return false
    }
}

// checks non empty
function validName(target) {
    let name = document.getElementById(target).value
    if (name == "") {
        return false
    }
    return true
}

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms))
}