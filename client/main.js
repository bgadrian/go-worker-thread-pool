
window.addEventListener('load', function () {
    console.log('All assets are loaded')
    updateWorker({ workerID: "111", status: "sssss" })

    Array
        .from(document.getElementsByClassName("job"))
        .forEach(function (element) {
            element.addEventListener('click', sendJob);
        });

    listenWorkers()
})

var updateWorker = function (p) {
    workerID = p.workerID
    status = p.status

    domStatus = document.getElementById("w-st-" + workerID)
    if (domStatus == null) {
        //we need to create it
        domWorker = document.createElement("div")
        domWorker.setAttribute("id", "w-" + workerID)
        domWorkers = document.getElementById("workers")
        domWorkers.appendChild(domWorker)

        domTitle = document.createElement("h5")
        domTitle.textContent = workerID
        domWorker.appendChild(domTitle)

        domStatus = document.createElement("span")
        domStatus.setAttribute("id", "w-st-" + workerID)
        domStatus.textContent = status
        domWorker.appendChild(domStatus)
    }
}

var sendJob = function () {
    magic = this.textContent
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

        console.log(msg.data)
    }
}