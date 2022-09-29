import Head from "next/head"
import ArticleCard from "../components/articlecard"
import SearchForm from "../components/searchbox"
import styles from "../styles/InfoPage.module.css"

export async function getServerSideProps(context){
    const searchquery = context.query
    const searchWord = searchquery.search

    const host = context.req.headers.host || 'localhost:3000'
    const protocol = /^localhost/.test(host) ? 'http' : 'https' 
    const notionData = await fetch(`${protocol}://${host}/api/search/notion/${searchWord}`)
        .then(data => data.json())
    
    return {
       props: {
        notionData,
        searchWord,
       },
    }
}

export default function infoPage({notionData, searchWord}){

    let articleCardList = []
    for(let i in notionData.results){
        articleCardList.push(
            <ArticleCard url={notionData.results[i].url} paragraph={notionData.results[i].context} key={i} />
        )
    }

    return (
        <>
            <Head>
                <title>PIPO</title>
            </Head>
            <div className={styles.infosearch}>
                <SearchForm initState={searchWord} />
            </div>
            <div className={styles.container}>
                {articleCardList}
            </div>
        </>
    )
}