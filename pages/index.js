import Head from 'next/head';
import SearchForm from '../components/searchbox';
import styles from '../styles/Home.module.css';

export default function Home(props) {
  
  return (
    <div className={styles.searchpage}>
      <Head>
        <title>PIPO</title>
      </Head>
      <SearchForm />
    </div>
  )
}