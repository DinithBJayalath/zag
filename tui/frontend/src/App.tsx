import {useState} from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import TerminalView from './TerminalView';
import {Greet} from "../wailsjs/go/main/App";

function App() {
    return (
        <div style={{ height: "100vh", width: "100vw" }}>
            <TerminalView />
        </div>
    );
}

export default App
