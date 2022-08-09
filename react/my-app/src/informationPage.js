import {React} from "react";
import {useLocation} from 'react-router-dom';
import SearchForm from "./search";
import ArticleCard from "./articleCard";

function InfoPage() {
    // stateを取り出す
    const location = useLocation();
    const prevSearchWords = String(location.state);

    var articleCardList = [
        <ArticleCard url="https://www.sejuku.net/blog/60444"/>,
        <ArticleCard />,
        <ArticleCard />,
        <ArticleCard /> ];

    return (
        <div className="Info-page">
            <header className="Info-header">
                <SearchForm initState={prevSearchWords} />
            </header>
            <body className="Info-body">
                <div className="Article-container">
                    {articleCardList}
                </div>
            </body>
        </div>
    );
}

export default InfoPage;