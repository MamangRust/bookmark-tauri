import React from 'react';

import CreateMarkdown from './pages/CreateMarkdown';
import AllMarkdown from './pages/AllMarkdown';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import FindMarkdown from './pages/FindMarkdown';
import AllFolder from './pages/AllFolder';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<AllFolder />} />
        <Route path="/find-file/:folder/:filename" element={<FindMarkdown />} />
        <Route path="/all-text/:folder" element={<AllMarkdown />} />
        <Route path="/create" element={<CreateMarkdown />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
