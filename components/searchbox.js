import Image from 'next/image';
import { useRouter } from 'next/router';
import {React, useState} from 'react';
import styles from '../styles/Search.module.css'

export default function SearchForm(props) {
    const router = useRouter()

    let [state, setState] = useState(props.initState);

    // とりまnull消すけどこれ辞めないと
    if (state==="null" || state==="undefined" || state===null || state===undefined) {
        setState("")
    }

    function handleChange(event) {
        setState(event.target.value);
    }

    function handleSubmit(event, props) {
        event.preventDefault();
        router.push({
        pathname: '/infopage',
        query: { search: state}
        })
    }
    return(
    <div className={styles.Search}>
        <form onSubmit={handleSubmit} >
          <div className={styles.img}>
            <Image src='/search_grass.png' layout='fill' objectFit='contain' alt="searchGrass" />
          </div>
          <input type="text" className={styles.input} placeholder="Pipoで検索" value={state} onChange={handleChange} minLength="1" />
        </form>
    </div>
    )
}

// propsのデフォルト値
SearchForm.defaultProps = {
    initState: "",
}