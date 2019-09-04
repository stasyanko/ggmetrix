import React, { Component } from 'react';
import axios from 'axios';

import ChartComponent from './ChartComponent';

export default class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      select_options: [],
      create_form: {
        //TODO: make array of constants for types of metrics
        'type': '',
        'title': '',
      },
      counter_data: [],
    };
  }

  async componentDidMount() {
    let selectOptsResponse = await axios(process.env.MIX_GO_API_URL + "/select_options");
    let countersResponse = await axios(process.env.MIX_GO_API_URL + "/counter/" + selectOptsResponse.data.data[0]['Title']);

    this.setState({
      select_options: selectOptsResponse.data.data,
      counter_data: countersResponse.data.data,
    });
  }

  async createMetricsType() {
    if (this.state.create_form.type === 'counter') {
      await axios.post(process.env.MIX_GO_API_URL + "/counter", {
        title: this.state.create_form.title
      });
      this.setState(prevState => ({
        select_options: [...prevState.select_options, {
          //TODO: replace with real ID
          ID: 5,
          Title: this.state.create_form.title,
          Type: this.state.create_form.type,
        }],
        create_form: {
          title: '',
          type: '',
        }
      }));
    }
  }

  async setCurType(e) {
    let a = e.target.value;
    let countersResponse = await axios(process.env.MIX_GO_API_URL + "/counter/" + e.target.value);
    debugger;
    this.setState({
      counter_data: countersResponse.data.data,
    });
  }

  render() {
    const selectOptionsReady = this.state.select_options.map(option => {
      return <option key={option.ID} value={option.Title}>{option.Title + "." + option.Type}</option>
    });

    return (
      <div class="row d-flex">
        <div className="col-sm-12 col-md-6 offset-md-3" style={{ paddingTop: '15px' }}>
          <select className="form-control" onChange={(e) => { this.setCurType(e) }}>
            {selectOptionsReady}
          </select>
        </div>

        <div className="col-sm-12 col-md-6 offset-md-3" style={{ paddingTop: '15px' }}>
          <ChartComponent data={this.state.counter_data} />
        </div>
      </div>
    );
  }
}

if (document.getElementById('react-root')) {
  ReactDOM.render(<App />, document.getElementById('react-root'));
}
