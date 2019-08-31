import React, { Component } from 'react';
// import axios from 'axios';

export default class App extends Component {
  constructor(props) {
    super(props);
    this.state = {

    };
  }

  componentDidMount() {
    // async function fetchMyAPI() {
    //   let response = await axios(
    //     'http://localhost:3000/select_options',
    //   )
    //   let res = response.data
    //   console.log(res);
    // }
  }

  render() {
    const selectOptionsReady = [].map(option => {
      return <option>{option}</option>
    });

    return (
      <div class="row d-flex">
        <div class="col-sm-12 col-md-6 offset-md-3" style={{ paddingTop: '15px' }}>
          <select class="form-control">
            {selectOptionsReady}
          </select>
        </div>

        <div class="col-sm-12 col-md-6 offset-md-3" style={{ paddingTop: '15px' }}>

        </div>
      </div>
    );
  }
}

if (document.getElementById('react-root')) {
  ReactDOM.render(<App />, document.getElementById('react-root'));
}
