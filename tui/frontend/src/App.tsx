import {useState} from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import TerminalView from './TerminalView';

function App() {
    return (
        <div style={{ height: "100vh", width: "100vw" }}>
            <TerminalView />
        </div>
    );
}

export default App
