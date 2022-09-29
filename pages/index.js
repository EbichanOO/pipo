import Head from 'next/head';
import Image from 'next/image';
import {useState} from 'react';
import SearchForm from '../components/searchbox';
import styles from '../styles/Home.module.css';

export default function Home(props) {
  const [isActiveNotion, setIsActiveNotion] = useState(false);
  const [isActive, setIsActive] = useState(false);

  const handleClickNotion = () => {
    setIsActiveNotion(current => !current);
  };
  const handleClick = () => {
    setIsActive(current => !current);
  };
  
  return (
    <div className={styles.searchpage}>
      <Head>
        <title>PIPO</title>
      </Head>
      <div className={styles.searchbox}>
        <SearchForm />
        <div className={styles.targetbox}>
          <div className={styles.notionbuttom}
          style={{
            border: isActiveNotion ? '' : 'thick solid black',
        }}
        onClick={handleClickNotion}>
            <Image src={'/icon-notion.png'} layout='fill' objectFit='contain' alt="notion"/>
          </div>
          <div className={styles.twitterbuttom}
          style={{
            border: isActive ? 'thick solid black' : '',
          }}
          onClick={handleClick}>
            <Image src={'/2021 Twitter logo - blue.png'} layout='fill' objectFit='contain' alt="notion"/>
          </div>
        </div>
      </div>
    </div>
  )
}