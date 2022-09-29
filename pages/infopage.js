import Head from "next/head"
import ArticleCard from "../components/articlecard"
import SearchForm from "../components/searchbox"
import styles from "../styles/InfoPage.module.css"

export async function getServerSideProps(context){
    const searchquery = context.query
    const searchWord = searchquery.search
    return {
       props: {searchWord},
    }
}

export default function infoPage({ searchWord }){
    var articleCardList = [
        <ArticleCard url="https://www.sejuku.net/blog/60444" key={"sample"}/>,
        <ArticleCard key={"1"}/>,
        <ArticleCard key={"2"}/>,
        <ArticleCard key={"3"}/> 
    ]

    return (
        <>
            <Head>
                <title>PIPO</title>
            </Head>
            <SearchForm initState={searchWord} />
            <div className={styles.container}>
                {articleCardList}
            </div>
        </>
    )
}