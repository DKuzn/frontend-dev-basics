async function onFormSubmit() {
    let form = document.getElementById("mainForm");

    let form_data = {
        event_name: form.elements["event_name"].value,
        event_plate: form.elements["event_plate"].value,
        event_type: form.elements["event_type"].value,
        event_date: form.elements["event_date"].value,
        event_time: form.elements["event_time"].value,
    };

    form_data.event_date = new Date(form_data.event_date).toLocaleDateString();

    fetch("http://localhost:8080/api/event", {
        method: "POST",
        mode: "cors",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(form_data)
    });
}

async function onLoadHandler() {
    let response = await fetch("http://localhost:8080/api/event", {
        method: "GET",
        mode: "cors",
        headers: {
            "Content-Type": "application/json"
        }
    });

    let ids = await response.json();

    formatTable(ids);
}

async function formatTable(ids) {
    let table = document.getElementById("dataTable");
    clearTable(table);

    for (let i = 0; i < ids.length; i++) {
        let response = await fetch("http://localhost:8080/api/event/" + ids[i], {
            method: "GET",
            mode: "cors",
            headers: {
                "Content-Type": "application/json"
            }
        });

        let data = await response.json();

        if (i === 0) {
            let headers = Object.keys(data);
            addHeaderToTable(table, headers);
        }

        addRowToTable(table, data);
    }
}

function clearTable(table) {
    table.innerHTML = "";
}

function addHeaderToTable(table, headers) {
    let headerElement = _createRow("th", headers);

    table.appendChild(headerElement);
}

function addRowToTable(table, row) {
    let data = Object.values(row);
    let rowElement = _createRow("td", data);

    table.appendChild(rowElement);
}

function _createRow(rowName, rowData) {
    let rowElement = document.createElement("tr");
    for (r of rowData) {
        let row = document.createElement(rowName)
        row.textContent = r;
        rowElement.appendChild(row);
    }

    return rowElement;
}