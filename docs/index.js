async function main(origin) {
    let response = await fetch(`${origin}-status.json`)
    let sites = await response.json()

    sitesIPV6 = sites.filter(item => item.ipv6)


    const elem = document.querySelector(`#${origin.toLowerCase()}-usage`)
    elem.textContent = sitesIPV6.length/sites.length*100

    const collection = document.querySelector(`#${origin.toLowerCase()}-details`)

    for (const site of sites) {
        const li = document.createElement('li')
        li.className = 'collection-item'

        if (site.ipv6) {
            li.textContent = site.domain + ' ✅'
        } else {
            li.textContent = site.domain + ' ❌'

        }

        collection.appendChild(li)
    }
}

main("Global")
main("BR")
