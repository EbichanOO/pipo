import { useRouter } from 'next/router';
import styles from '../styles/ArticleCard.module.css';

export default function ArticleCard(props) {
    const router = useRouter()
    const articleURL = props.url;
    const articleTitle = props.title;
    const articleParagraph = props.paragraph;
    function redirectToUrl() {router.push(articleURL)}
    return (
        <div className={styles.card} onClick={redirectToUrl}>
            <h3>{articleTitle}</h3>
            <p>{articleParagraph}</p>
        </div>
    );
}

ArticleCard.defaultProps = {
    url: "https://style.potepan.com/articles/21691.html",
    title: "",
    paragraph: "",
};