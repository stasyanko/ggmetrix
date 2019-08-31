import React, { Component } from 'react';
import axios from 'axios';

export default class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      select_options: []
    };
  }

  async componentDidMount() {
    let response = await axios(process.env.MIX_GO_API_URL + "/select_options");
    this.setState({
      select_options: response.data.data
    });
  }

  render() {
    const selectOptionsReady = this.state.select_options.map(option => {
      return <option key={option.ID}>{option.Title + "." + option.Type}</option>
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
