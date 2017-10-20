import React from 'react';
import ReactDOM from 'react-dom';
import FormWidget from './components/FormWidget.jsx';

import lightBaseTheme from 'material-ui/styles/baseThemes/lightBaseTheme';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import getMuiTheme from 'material-ui/styles/getMuiTheme';

// import '../styles/application.scss';

const Application = () => (
    <MuiThemeProvider muiTheme={getMuiTheme(lightBaseTheme)}>
        <FormWidget />
    </MuiThemeProvider>
);

document.addEventListener("DOMContentLoaded", () => {
    ReactDOM.render(<Application />, document.getElementById('react-application'));
});