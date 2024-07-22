import React, { useState } from 'react';
import { setCache, getCache, deleteCache } from './Api';

function App() {
  const [key, setKey] = useState('');
  const [value, setValue] = useState('');
  const [expiration, setExpiration] = useState('5s');
  const [result, setResult] = useState('');

  const handleSet = async () => {
    await setCache(key, value, expiration);
    setResult('Cache set successfully');
  };

  const handleGet = async () => {
    const value = await getCache(key);
    setResult(value !== null ? value : 'Key not found');
  };

  const handleDelete = async () => {
    await deleteCache(key);
    setResult('Cache deleted successfully');
  };

  return (
    <div>
      <h1>LRU Cache</h1>
      <input
        type="text"
        value={key}
        onChange={(e) => setKey(e.target.value)}
        placeholder="Key"
      />
      <input
        type="text"
        value={value}
        onChange={(e) => setValue(e.target.value)}
        placeholder="Value"
      />
      <input
        type="text"
        value={expiration}
        onChange={(e) => setExpiration(e.target.value)}
        placeholder="Expiration (e.g., 5s)"
      />
      <button onClick={handleSet}>Set</button>
      <button onClick={handleGet}>Get</button>
      <button onClick={handleDelete}>Delete</button>
      <p>Result: {result}</p>
    </div>
  );
}

export default App;
