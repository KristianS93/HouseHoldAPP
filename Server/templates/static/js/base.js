async function login() {
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
            // login request was acknowledged as successful
            location.reload()
            break
        case 404:
            // user credentials did not match anything in database
            
            // should generate a pop up alert inside the modal
            // should be Warning alert
            break
        default:
            // should essentially only consist of 500 status code
            // should make Danger alert
            break
    }
}

// currently only updates login button based on focus changing from one of the text input fields
// possibly find another solution, or instead remove disable part and instead show a warning if login is tried without acceptable input
function disableLogin() {
    let btn = document.getElementById("loginButton")
    if (validEmail() && validPassword()) {
        btn.disabled = false
    } else {
        btn.disabled = true
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