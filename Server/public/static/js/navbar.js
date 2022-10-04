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

    let response = await fetch("http://localhost:8888/register", 
    {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            FirstName: document.getElementById(Elements.RegisterFirstName).value,
            LastName: document.getElementById(Elements.RegisterLastName).value,
            UserID: document.getElementById(Elements.RegisterEmail).value,
            Password: document.getElementById(Elements.RegisterPassword).value,
            ConfirmPassword: document.getElementById(Elements.RegisterPasswordConfirm)
        })
    })

    switch (response.status) {
        case HTTPCode.Success:
            // should read the session cookie sent from registration backend
            // and set it serverside, so that a reload will login the user (as session is already active)
            addAlert("success", alertParagraph("Registration was successful."), Elements.RegisterModalBody)
            await sleep(1000)
            location.reload()
            break
        case HTTPCode.BadRequest:
            console.log("bad request")
            let body = await response.json()
            let concatErr = ""
            for(let error of body.Errors) {
                concatErr += alertParagraph(error)
            }
            addAlert("warning", concatErr, Elements.RegisterModalBody)
            break
        case HTTPCode.Conflict:
            addAlert("warning", alertParagraph("A user is already registered with that email."), Elements.RegisterModalBody)
            break
        case HTTPCode.InternalServerError:
            addAlert("danger", alertParagraph("Internal server error."), Elements.RegisterModalBody)
            break
        default:
            addAlert("danger", alertParagraph("Unknown error occured."), Elements.RegisterModalBody)
            console.log("Register attempt status code: " + response.status)
            break
    }
}

function checkRegistration() {
    let returnString = ""
    if (!validName(Elements.RegisterFirstName)) {
        returnString += alertParagraph("First Name is empty.")
    }
    if (!validName(Elements.RegisterLastName)) {
        returnString += alertParagraph("Last Name is empty.")
    }
    if (!validEmail(Elements.RegisterEmail)) {
        returnString += alertParagraph("Email format is invalid.")
    }
    if (document.getElementById(Elements.RegisterPassword).value != document.getElementById(Elements.RegisterPasswordConfirm).value) {
        returnString += alertParagraph("Passwords do not match.")
    }
    if (!validPassword(Elements.RegisterPassword)) {
        returnString += alertParagraph("Password is not an appropriate length.")
    }
    return returnString
}

function alertParagraph(message) {
    return `<p class="mb-0">${message}</p>`
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
            addAlert("success", alertParagraph("Login was successful."), Elements.LoginModalBody)
            await sleep(1000)
            location.reload()
            break
        case HTTPCode.BadRequest:
            addAlert("warning", alertParagraph("The provided login information are invalid. (S)"), Elements.LoginModalBody)
            console.log("validation: server")
            break
        case HTTPCode.NotFound:
            addAlert("warning", alertParagraph("No user was found with provided credentials."), Elements.LoginModalBody)
            break
        case HTTPCode.InternalServerError:
            addAlert("danger", alertParagraph("Internal server error."), Elements.LoginModalBody)
            break
        default:
            addAlert("danger", alertParagraph("Unknown error occured."), Elements.LoginModalBody)
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