import Image from 'next/image';
import Head from 'next/head';
import { useRouter } from 'next/router';
import {React, useState} from 'react';
import styles from '../styles/Search.module.css';
import SearchForm from '../components/searchbox';

export default function SearchPage(props) {
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
  
  return (
    <>
      <Head>
        <title>PIPO SEARCH PAGE</title>
      </Head>
      <SearchForm />
    </>
  )
}

// propsのデフォルト値
SearchPage.defaultProps = {
  initState: "",
}