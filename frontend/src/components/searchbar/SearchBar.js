import React, { useState } from 'react';
import "../../css/leftmenu.css"

function SearchBar({ onSearch }) {
  const [query, setQuery] = useState('');
  const handleSearch = (event) => {
    const newQuery = event.target.value;
    setQuery(newQuery);
    onSearch(newQuery);
  };
  return (
    <div className="search-bar">
  <input id='search' type="text" value={query} onChange={handleSearch} placeholder="Rechercher..." />
</div>
  );
}
export default SearchBar;