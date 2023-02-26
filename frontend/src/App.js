import cl from './styles/App.css'
import React, {useState,userRef} from 'react';
import { BrowserRouter, Route, Routes} from 'react-router-dom';
import Start from './Pages/Start';
import Tasks from './Pages/Tasks';

function App() {
  return (
    <BrowserRouter>
        <Routes>
      <Route path="/" element={<Start />} />
      <Route path="Tasks" element={<Tasks />} />
    </Routes>
    </BrowserRouter>
  )
}

export default App;
