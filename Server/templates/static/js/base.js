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

var RegEx = {
    LowerCase: /[a-z]/,
    UpperCase: /[A-Z]/,
    Number: /[0-9]/,
    Special: /[!@#$%^&*]/,
    EmailIdentifier: /[a-zA-Z._-]/g,
    Names: /[a-zA-Z'-]/g,
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
        returnString += listStart+"First Name does not fit required format."+listEnd
    }
    if (validName(Elements.RegisterLastName)) {
        returnString += listStart+"Last Name does not fit required format."+listEnd
    }
    if (!validEmail(Elements.RegisterEmail)) {
        returnString += listStart+"Email format is invalid."+listEnd
    }
    if (document.getElementById(Elements.RegisterPassword) != document.getElementById(Elements.RegisterPasswordConfirm)) {
        returnString += listStart+"Passwords do not match."+listEnd
    }
    if (!validPassword(Elements.RegisterPassword)) {
        returnString += listStart+"Password does not match required criteria."+listEnd
    }
    return returnString
}

// this needs to be fixed
function validName(name) {
    if (name.match(RegEx.Names)) {
        return true
    }
    return false
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
            document.cookie = "success=Login was successful.; max-age=2"
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

function validEmail(targetID) {
    let email = document.getElementById(targetID).value
    let splitEmail = email.split("@")

    if (splitEmail.length != 2) {
        return false
    }

    if (splitEmail[0].match(RegEx.EmailIdentifier) && splitEmail[1].match(/[.]/)) {
        return true
    }
    return false
}

function validPassword(targetID) {
    let password = document.getElementById(targetID).value
    let hasLength = false, hasUpper = false, hasLower = false, hasNumber = false, hasSpecial = false

    if (password.length >= 8 && password.length <= 32) {
        hasLength = true
    } else {
        return false
    }

    for (const char of password) {
        if (RegEx.LowerCase.test(char)) {
            hasLower = true
        }
        if (RegEx.UpperCase.test(char)) {
            hasUpper = true
        }
        if (RegEx.Number.test(char)) {
            hasNumber = true
        }
        if (RegEx.Special.test(char)) {
            hasSpecial = true
        }
    }
    return hasLength && hasUpper && hasLower && hasNumber && hasSpecial
}

// test cases for password
// console.log("aaaaaaaaaaa " + password("aaaaaaaaaaa"))
// console.log("AAAAAAAAAAA " + password("AAAAAAAAAAA"))
// console.log("11111111111 " + password("11111111111"))
// console.log("########### " + password("###########"))
// console.log("PEYAPsPukyYj$pfY4D^3m$x$49*p2Tt8 " + password("PEYAPsPukyYj$pfY4D^3m$x$49*p2Tt8"))
// console.log("r8kq4dmxj*y%b3xtv$@b953&r@i$!o^u " + password("r8kq4dmxj*y%b3xtv$@b953&r@i$!o^u"))
// console.log("2TKE5U%GZCC3^62S7UW*U%UKYQZBCMA8 " + password("2TKE5U%GZCC3^62S7UW*U%UKYQZBCMA8"))
// console.log("sQECZ*cgacDZ&*iqPYqhHwg@fogQb*R^ " + password("sQECZ*cgacDZ&*iqPYqhHwg@fogQb*R^"))
// console.log("HQsbYxfGgYw4FT3t6fKAhS93YPqHpeFT " + password("HQsbYxfGgYw4FT3t6fKAhS93YPqHpeFT"))