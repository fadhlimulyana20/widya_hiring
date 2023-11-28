export function buildQuery(param: Object) {
    let query = []
    for (const [k, v] of Object.entries(param)) {
        if (v !== "" && v !== null ) {
            if (typeof v !== 'undefined') {
                query.push(`${k}=${encodeURIComponent(v)}`)
            }
        }
    }
    return query.join("&")
}
