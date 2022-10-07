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
function toggleRead(id) {
    let list = document.querySelectorAll(`#${id}`);
    for (let i = 0; i < list.length; i++) {
        let element = list[i];
        if (element.readOnly === true) {
            element.readOnly = false;
        }
        else {
            element.readOnly = true;
        }
    }
    if (list[0].readOnly === false) {
        toggleButtons(id);
    }
    else {
        toggleButtons(id);
    }
}
function toggleButtons(id) {
    let btn = document.getElementById(id + 0);
    if (btn.disabled === true) {
        btn.disabled = false;
    }
    else {
        btn.disabled = true;
    }
    btn = document.getElementById(id + 1);
    if (btn.disabled === true) {
        btn.disabled = false;
    }
    else {
        btn.disabled = true;
    }
}
function changeItem(id) {
    return __awaiter(this, void 0, void 0, function* () {
        toggleRead(id);
        let list = document.querySelectorAll(`#${id}`);
        let name = list[0].value, quantity = list[1].value, unit = list[2].value;
        console.log(name, quantity, unit);
        let Item = {
            Id: id,
            ItemName: name,
            Quantity: quantity,
            Unit: unit
        };
        let response = yield fetch("http://localhost:8888/changeitem", {
            method: 'PATCH',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(Item)
        });
        if (response.status != 200) {
            location.reload();
        }
    });
}
