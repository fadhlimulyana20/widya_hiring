import { useEffect, useRef, useState } from "react";

declare global {
    interface Window {
        MathJax: any
    }
}


export function Latex({ text }: { text: string }) {
    const refNode = useRef<HTMLDivElement>(null)
    const [tex, setTex] = useState('...')

    const typeset = (code: any) => {
        let promise = Promise.resolve();
        promise = promise.then(() => window.MathJax.typesetPromise(code()))
            .catch((err) => console.log('Typeset failed: ' + err.message));
        return promise;
    }


    useEffect(() => {
        // console.log('update')
        setTex(text)
        typeset(() => {
            // let d = document.querySelector('#_mathjaxdiv')
            return [refNode.current]
        }
        )
    }, [tex])

    return (
        <>
            <div ref={refNode} id="_mathjaxdiv">
                {tex}
            </div>
        </>
    )
}

function splitHTML(html: string) {
    const regex = /<p>(.*?)<\/p>|<h1>(.*?)<\/h1>|<h2>(.*?)<\/h2>|<h3>(.*?)<\/h3>|<h4>(.*?)<\/h4>|<h5>(.*?)<\/h5>|<ul>(.*)<\/ul>|<ol>(.*?)<\/ol>/
    let htmls = html.match(regex)
    return htmls
}

export function LatexHTML({ html }: { html: string }) {
    const refNode = useRef<HTMLDivElement>(null)
    const [htmlElements, setHtmlElements] = useState<RegExpMatchArray>()

    useEffect(() => {
      const htmls = splitHTML(html)
      if (htmls !== null) {
          setHtmlElements(htmls)
      }

    }, [])


    return (
        <>
            <div>
                { htmlElements?.map((obj, ix) => {
                     return <div key={ix} dangerouslySetInnerHTML={{__html: obj}} />
                    // if (obj.startsWith("$")) {
                    //     return <Latex text={obj} />
                    // } else {
                    //     return <div key={ix} dangerouslySetInnerHTML={{__html: obj}} />
                    // }
                }) }
            </div>
        </>
    )
}
