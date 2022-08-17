function toggleRead(id) {
    let list = document.querySelectorAll(`#${id}`)
    for (let i = 0; i < list.length; i++) {
        if (list[i].readOnly === true) {
            list[i].readOnly = false
        } else {
            list[i].readOnly = true
        }
    }
    if (list[0].readOnly === false) {
        toggleDisable(id)

    } else {
        toggleDisable(id)
    }
}

function toggleDisable(id) {
    let btn = document.getElementById(id+0)
    if (btn.disabled === true) {
        btn.disabled = false
    } else {
        btn.disabled = true
    }
}

async function changeItem(id) {
    toggleRead(id)
    let list = document.querySelectorAll(`#${id}`)
    let name = list[0].value, quantity = list[1].value, unit = list[2].value
    console.log(name, quantity, unit)

    let Item = {
        Id: id,
        Name: name,
        Quantity: quantity,
        Unit: unit
    }
    console.log(Item)

    let response = await fetch("http://localhost:8888/changeitem", 
    {
        method: 'PATCH',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(Item)
    })

    if (response.status != 200) {
        location.reload()
    }
}