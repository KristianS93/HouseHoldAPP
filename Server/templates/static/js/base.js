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

async function login() {
    if (!checkLogin()) {
        addAlert("warning", "Invalid login credentials.", "loginModalBody")
        return
    }

    // not sure if this is bad practice or not lmao
    let user = {
        UserID: document.getElementById("loginEmail"),
        Password: document.getElementById("loginPassword")
    }

    let response = await fetch("http://localhost:8888/login", 
    {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(user)
    })

    switch (response.status) {
        case 200:
            // the below does not work currently, unsure of why or how - it does not send cookies to 
            document.cookie = "success=Login was successful.; max-age=2"
            location.reload()
            break
        case 404:
            addAlert("warning", "No user was found with provided credentials.", "loginModalBody")
            break
        default:
            // should essentially only consist of 500 status code
            // should make Danger alert
            addAlert("danger", "Internal server error.", "loginModalBody")
            break
    }
}

// currently only updates login button based on focus changing from one of the text input fields
// possibly find another solution, or instead remove disable part and instead show a warning if login is tried without acceptable input
function checkLogin() {
    if (validEmail() && validPassword()) {
        return true
    } else {
        return false
    }
}

function validEmail() {
    let regEx = /^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$/
    if (document.getElementById("loginEmail").value.match(regEx)) {
        return true
    } else {
        return false
    }
}

function validPassword() {
    let l = document.getElementById("loginPassword").value.length
    if (l >= 8 && l <= 32) {
        return true
    } else {
        return false
    }
}

function disableRegister() {
    // same as above, just for register
    // should also have helper to check password length, if something is an email, and password + repassword match
}

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