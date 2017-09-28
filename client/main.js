var globalCounter = 0

window.addEventListener('load', function () {
    console.log('All assets are loaded')
    // updateWorker({ workerID: "111", status: "sssss" })

    Array
        .from(document.getElementsByClassName("job"))
        .forEach(function (element) {
            element.addEventListener('click', sendJob);
        });

    listenWorkers()
})

var updateWorker = function (p) {
    var workerID = p["WorkerID"]
    var status = p["Status"]

    var statusID = "w-st-" + workerID
    var domStatus = document.getElementById(statusID)

    console.log("updateWorker", JSON.stringify(p), statusID, domStatus)
    if (domStatus == null) {
        //we need to create it
        var domWorker = document.createElement("div")
        domWorker.setAttribute("id", "w-" + workerID)
        domWorkers = document.getElementById("workers")
        domWorkers.appendChild(domWorker)

        var domTitle = document.createElement("h5")
        domTitle.textContent = workerID
        domWorker.appendChild(domTitle)

        var domStatus = document.createElement("span")
        domStatus.setAttribute("id", "w-st-" + workerID)
        domWorker.appendChild(domStatus)
    }

    domStatus.textContent = status
}

var sendJob = function () {
    magic = this.textContent
    magic += "[" + (globalCounter++) + "]"
    request = new Request("/job",
        {
            method: 'POST',
            body: '{"Magic":"' + magic + '"}'
        });


    fetch(request)
        .catch(function (error) {
            console.error(error);
        });
}

var listenWorkers = function () {
    socket = new WebSocket("ws://localhost:8080/ws")

    socket.onopen = function () {
        console.log("websocket open")
    }

    socket.onclose = function () {
        console.log("websocket close")
    }

    socket.onmessage = function (msg) {
        json = JSON.parse(msg.data)

        updateWorker(json)
    }
}