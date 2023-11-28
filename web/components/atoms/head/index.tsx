import Head from "next/head";
import { useRouter } from "next/router";
import { useEffect, useState } from "react";

interface props {
    title: string;
    description: string;
    ogTitle?: string;
    ogType?: string;
    ogUrl?: string;
    ogImage?: string;
    keywords?: string;
}

export default function Header({
    title,
    description,
    ogTitle="",
    ogType="page",
    ogUrl="",
    ogImage="",
    keywords=""
} : props) {
    const router = useRouter()
    const [suffixTitle, setSuffixTitle] = useState("")



    return (
        <Head>
            <title>{title + suffixTitle}</title>
            <meta name="description" content={description}></meta>
            <meta name="og:description" content={description}></meta>
            <meta name="twitter:description" content={description}></meta>
            <meta property="og:title" content={ogTitle === "" ? title : ogTitle} />
            <meta property="og:type" content={ ogType } />
            <meta property="og:url" content={ ogUrl} />
            <meta property="og:image" content={ ogImage } />
            <meta property="og:site_name" content="" />
            <meta name="keywords" content={ keywords }></meta>
        </Head>
    )
}
