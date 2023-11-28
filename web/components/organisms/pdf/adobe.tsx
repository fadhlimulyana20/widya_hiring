import Script from "next/script"
import { useRef } from "react"

export function AdobePDFViewer() {
    const divRef = useRef<HTMLDivElement>(null)

    return (
        <>
            <Script src="https://acrobatservices.adobe.com/view-sdk/viewer.js" />
            <div ref={divRef} />
        </>
    )
}
