import Head from 'next/head';
import SearchForm from '../components/searchbox';

export default function Home(props) {
  
  return (
    <>
      <Head>
        <title>PIPO</title>
      </Head>
      <SearchForm />
    </>
  )
}