import { BrowserRouter, Route, Routes } from "react-router-dom";
import './App.css';
import Header from './component/header/Header.tsx';
import ThreadCreation from './component/pages/ThreadCreation.tsx';

function App() {
  return (
    <div>
      <Header></Header>
      sssss
      <BrowserRouter>
        <Routes>
          {/* <Route path="/" element={<Home />} /> */}
          <Route path="/thread/create" element={<ThreadCreation />} />
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
