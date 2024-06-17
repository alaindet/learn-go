import { useEffect, useState } from 'react';
import './App.css';

const fetchData = async (url: string) => {
  return fetch(url).then(res => res.json());
};

export function App() {

  const [serverMessage, setServerMessage] = useState('Pending...');

  async function fetchServerMessage() {
    const url = 'http://localhost:8080/api';
    const res = await fetchData(url);
    const message = res['data'];
    setServerMessage(message);
  }

  useEffect(() => {
    fetchServerMessage();
  }, []);

  return (
    <div className="app">
      <p>This is a demo React app embedded into a Go binary</p>
      <p>The server said: {serverMessage}</p>
    </div>
  ); 
}
