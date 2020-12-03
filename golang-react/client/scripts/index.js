import React from 'react';
import ReactDom from 'react-dom';
import App from './main/App';
import Config from 'Config';

const app = document.getElementById('app');

ReactDom.render(<App serverURL={Config.serverURL}/>, app);
