function toggleRead(id: string) {
    let list = document.querySelectorAll(`#${id}`)
    


    for (let i: number = 0; i < list.length; i++) {
        let element = (<HTMLInputElement>list[i])
        if (element.readOnly === true) {
            element.readOnly = false
        } else {
            element.readOnly = true
        }
    }

    if ((<HTMLInputElement>list[0]).readOnly === false) {
        toggleButtons(id)
    } else {
        toggleButtons(id)
    }
}

function toggleButtons(id: string) {
    let btn = (<HTMLButtonElement>document.getElementById(id+0))
    if (btn.disabled === true) {
        btn.disabled = false
    } else {
        btn.disabled = true
    }

    btn = (<HTMLButtonElement>document.getElementById(id+1))
    if (btn.disabled === true) {
        btn.disabled = false
    } else {
        btn.disabled = true
    }
}

async function changeItem(id: string) {
    toggleRead(id)
    let list = document.querySelectorAll(`#${id}`)
    let name = (<HTMLInputElement>list[0]).value, quantity = (<HTMLInputElement>list[1]).value, unit = (<HTMLInputElement>list[2]).value
    console.log(name, quantity, unit)

    let Item = {
        Id: id,
        ItemName: name,
        Quantity: quantity,
        Unit: unit
    }

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