var HTTPCode = {
    Success: 200,
    BadRequest: 400,
    NotFound: 404,
    Conflict: 409,
    InternalServerError: 500,
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

async function login() {
    if (!checkLogin()) {
        addAlert("warning", "Invalid login credentials.", "loginModalBody")
        return
    }

    let response = await fetch("http://localhost:8888/login", 
    {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            UserID: document.getElementById("loginEmail").value,
            Password: document.getElementById("loginPassword").value
        })
    })

    switch (response.status) {
        case HTTPCode.Success:
            document.cookie = "success=Login was successful.; max-age=2"
            location.reload()
            break
        case HTTPCode.BadRequest:
            addAlert("warning", "The provided login information are invalid.", "loginModalBody")
            break
        case HTTPCode.NotFound:
            addAlert("warning", "No user was found with provided credentials.", "loginModalBody")
            break
        case HTTPCode.InternalServerError:
            addAlert("danger", "Internal server error.", "loginModalBody")
            break
        default:
            addAlert("danger", "Unknown error occured.", "loginModalBody")
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

// should remake to fit backend more
// function validEmail() {
//     let regEx = /^^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`/
//     if (document.getElementById("loginEmail").value.match(regEx)) {
//         return true
//     } else {
//         return false
//     }
// }

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

var RegEx = {
    LowerCase: /[a-z]/,
    UpperCase: /[A-Z]/,
    Number: /[0-9]/,
    Special: /[!@#$%^&*]/
  }
  
function validPassword() {
    let hasLength = false, hasUpper = false, hasLower = false, hasNumber = false, hasSpecial = false

    let p = document.getElementById("loginPassword").value
    if (p.length >= 8 && p.length <= 32) {
        hasLength = true
    } else {
        return false
    }

    for (const char of p) {
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