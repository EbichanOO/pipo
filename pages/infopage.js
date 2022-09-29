import Head from "next/head"
import ArticleCard from "../components/articlecard"
import SearchForm from "../components/searchbox"
import styles from "../styles/infoPage.module.css"

export async function getServerSideProps(context){
    const searchquery = context.query
    const searchWord = searchquery.search
    return {
       props: {searchWord},
    }
}

export default function infoPage({ searchWord }){
    var articleCardList = [
        <ArticleCard url="https://www.sejuku.net/blog/60444"/>,
        <ArticleCard />,
        <ArticleCard />,
        <ArticleCard /> 
    ]

    return (
        <>
            <Head>
                <title>hoge</title>
            </Head>
            <SearchForm initState={searchWord} />
            <div className={styles.container}>
                {articleCardList}
            </div>
        </>
    )
}