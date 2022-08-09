import {useNavigate} from 'react-router-dom';

function ArticleCard(props) {
    const articleURL = props.url;
    const articleTitle = props.title;
    const articleParagraph = props.paragraph;
    function redirectToUrl() {window.location.href = articleURL}
    return (
        <div className="Article-card" onClick={redirectToUrl}>
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

export default ArticleCard;