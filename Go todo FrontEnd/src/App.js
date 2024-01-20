import {Route, Routes} from 'react-router-dom';
import './App.css';
import Header from './components/Header/Header';
import Footer from './components/Footer/Footer';
import TodoPage from './pages/Todo/TodoPage';
import LoginPage from './pages/Login/LoginPage';
import RegisterPage from './pages/Register/Register';

function App() {
    return (
        <div className="App">
            <Header/>

            <Routes>
                <Route path='/' element={<TodoPage/>}/>
                <Route path='/login' element={<LoginPage/>}/>
                <Route path='/register' element={<RegisterPage/>}/>
            </Routes>

            <Footer/>
        </div>
    );
}

export default App;
