import React from 'react';
import MainPage from '../view/MainPage';

class App extends React.Component {
  render() {
    return <MainPage serverURL={this.props.serverURL}/>
  }
}

export default App;
