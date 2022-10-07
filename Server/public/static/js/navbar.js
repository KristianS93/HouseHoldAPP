"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var HTTPCode = {
    Success: 200,
    BadRequest: 400,
    NotFound: 404,
    Conflict: 409,
    InternalServerError: 500,
};
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
};
let x = document.getElementById("loginEmail");
x === null || x === void 0 ? void 0 : x.addEventListener("keydown", function (e) {
    if (e.code == "Enter") {
        login();
    }
});
let y = document.getElementById("loginPassword");
y === null || y === void 0 ? void 0 : y.addEventListener("keydown", function (e) {
    if (e.code == "Enter") {
        login();
    }
});
function addAlert(alertLevel, alertMessage, target) {
    let alert = `
    <div class="container">
      <div class="alert alert-${alertLevel} alert-dismissible fade show mx-auto" role="alert">
        ${alertMessage}
        <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>
      </div>
    </div>
    `;
    let element = document.getElementById(target);
    element === null || element === void 0 ? void 0 : element.insertAdjacentHTML("afterbegin", alert);
}
function register() {
    return __awaiter(this, void 0, void 0, function* () {
        let errorMessage = checkRegistration();
        if (errorMessage != "") {
            addAlert("warning", errorMessage, Elements.RegisterModalBody);
            return;
        }
        let response = yield fetch("http://localhost:8888/register", {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                FirstName: document.getElementById(Elements.RegisterFirstName).value,
                LastName: document.getElementById(Elements.RegisterLastName).value,
                UserID: document.getElementById(Elements.RegisterEmail).value,
                Password: document.getElementById(Elements.RegisterPassword).value,
                ConfirmPassword: document.getElementById(Elements.RegisterPasswordConfirm).value
            })
        });
        switch (response.status) {
            case HTTPCode.Success:
                // should read the session cookie sent from registration backend
                // and set it serverside, so that a reload will login the user (as session is already active)
                addAlert("success", alertParagraph("Registration was successful."), Elements.RegisterModalBody);
                yield sleep(1000);
                location.reload();
                break;
            case HTTPCode.BadRequest:
                console.log("bad request");
                let body = yield response.json();
                let concatErr = "";
                for (let error of body.Errors) {
                    concatErr += alertParagraph(error);
                }
                addAlert("warning", concatErr, Elements.RegisterModalBody);
                break;
            case HTTPCode.Conflict:
                addAlert("warning", alertParagraph("A user is already registered with that email."), Elements.RegisterModalBody);
                break;
            case HTTPCode.InternalServerError:
                addAlert("danger", alertParagraph("Internal server error."), Elements.RegisterModalBody);
                break;
            default:
                addAlert("danger", alertParagraph("Unknown error occured."), Elements.RegisterModalBody);
                console.log("Register attempt status code: " + response.status);
                break;
        }
    });
}
function checkRegistration() {
    let returnString = "";
    if (!validName(Elements.RegisterFirstName)) {
        returnString += alertParagraph("First Name is empty.");
    }
    if (!validName(Elements.RegisterLastName)) {
        returnString += alertParagraph("Last Name is empty.");
    }
    if (!validEmail(Elements.RegisterEmail)) {
        returnString += alertParagraph("Email format is invalid.");
    }
    if (document.getElementById(Elements.RegisterPassword).value != document.getElementById(Elements.RegisterPasswordConfirm).value) {
        returnString += alertParagraph("Passwords do not match.");
    }
    if (!validPassword(Elements.RegisterPassword)) {
        returnString += alertParagraph("Password is not an appropriate length.");
    }
    return returnString;
}
function alertParagraph(message) {
    return `<p class="mb-0">${message}</p>`;
}
function login() {
    return __awaiter(this, void 0, void 0, function* () {
        if (!checkLogin()) {
            addAlert("warning", "The provided login information is invalid.", Elements.LoginModalBody);
            console.log("validation: client");
            return;
        }
        let response = yield fetch("http://localhost:8888/login", {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                UserID: document.getElementById(Elements.LoginEmail).value,
                Password: document.getElementById(Elements.LoginPassword).value
            })
        });
        switch (response.status) {
            case HTTPCode.Success:
                addAlert("success", alertParagraph("Login was successful."), Elements.LoginModalBody);
                yield sleep(1000);
                location.reload();
                break;
            case HTTPCode.BadRequest:
                addAlert("warning", alertParagraph("The provided login information are invalid. (S)"), Elements.LoginModalBody);
                console.log("validation: server");
                break;
            case HTTPCode.NotFound:
                addAlert("warning", alertParagraph("No user was found with provided credentials."), Elements.LoginModalBody);
                break;
            case HTTPCode.InternalServerError:
                addAlert("danger", alertParagraph("Internal server error."), Elements.LoginModalBody);
                break;
            default:
                addAlert("danger", alertParagraph("Unknown error occured."), Elements.LoginModalBody);
                console.log("Login attempt status code: " + response.status);
                break;
        }
    });
}
// currently only updates login button based on focus changing from one of the text input fields
// possibly find another solution, or instead remove disable part and instead show a warning if login is tried without acceptable input
function checkLogin() {
    if (validEmail(Elements.LoginEmail) && validPassword(Elements.LoginPassword)) {
        return true;
    }
    else {
        return false;
    }
}
// only checks for @ and non empty
function validEmail(targetID) {
    let email = document.getElementById(targetID).value;
    if (email == "" || !email.includes("@")) {
        return false;
    }
    return true;
}
// only checkcs length
function validPassword(targetID) {
    let password = document.getElementById(targetID).value;
    if (password.length >= 8 && password.length <= 32) {
        return true;
    }
    else {
        return false;
    }
}
// checks non empty
function validName(target) {
    let name = document.getElementById(target).value;
    if (name == "") {
        return false;
    }
    return true;
}
function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}
