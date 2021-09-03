(async () => {
    const res = await fetch("/api/dir")
    const json = await res.json()
    for (const d of json.dir) {
        const hyperLink = document.createElement("a")
        hyperLink.href = "/" + d.name
        hyperLink.innerText = d.name
        const div = document.querySelector("div")
        div.appendChild(hyperLink)
        div.appendChild(document.createElement("br"))
    }
})()