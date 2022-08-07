import {React} from "react";
import {useLocation} from 'react-router-dom';
import SearchForm from "./search";
import ArticleCard from "./articleCard";

function InfoPage() {
    // stateを取り出す
    const location = useLocation();
    const prevSearchWords = String(location.state);

    return (
        <div className="Info-page">
            <body className="Info-body">
                <SearchForm initState={prevSearchWords} />
                <ArticleCard />
            </body>
        </div>
    );
}

export default InfoPage;