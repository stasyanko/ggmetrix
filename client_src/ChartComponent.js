import React, { PureComponent } from 'react';
import {
    ResponsiveContainer, ComposedChart, Line, Area, Bar, XAxis, YAxis, CartesianGrid, Tooltip, Legend,
} from 'recharts';
import moment from 'moment';

export default class ChartComponent extends PureComponent {
    constructor(props) {
        super(props);
    }

    render() {
        return (
            <div style={{ width: '100%', height: 300 }}>
                <ResponsiveContainer>
                    <ComposedChart
                        width={500}
                        height={400}
                        data={this.props.data}
                        margin={{
                            top: 20, right: 20, bottom: 20, left: 20,
                        }}
                    >
                        <CartesianGrid stroke="#f5f5f5" />
                        <XAxis type="number" name='Time' domain={['auto', 'auto']} dataKey="UnixTs" tickFormatter={unixTime => moment.unix(unixTime).format('DD/MM HH:mm')} />
                        <YAxis />
                        <Tooltip />
                        <Legend />
                        <Bar dataKey="Value" barSize={20} fill="#413ea0" />
                    </ComposedChart>
                </ResponsiveContainer>
            </div>
        );
    }
}
