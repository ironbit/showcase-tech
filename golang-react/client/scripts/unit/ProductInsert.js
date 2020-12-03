import React from 'react';
import Grid from '@material-ui/core/Grid';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import { withStyles } from '@material-ui/core/styles';

const styles = theme => ({
  textField: {
    //marginLeft: theme.spacing(2),
    //marginRight: theme.spacing(2),
    width: '40ch',
  },
  button: {
    size: 'small',
  }
});

class ProductInsert extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      product: {
        id: "",
        name: ""
      }
    };

    this.onButtonClick = this.onButtonClick.bind(this);
    this.handleTextFieldChange = this.handleTextFieldChange.bind(this);
  }

  onButtonClick() {
    this.props.onProductInsert(this.state.product)
    this.setState({
      product: {
        id: "",
        name: ""
      }
    })
  }

  handleTextFieldChange(event) {
    const target = event.target;
    if (target.name == 'input-id') {
      this.setState((state) => {
        let id = parseInt(target.value)
        if (isNaN(id)) {
          id = ""
        }
        return {
          product: {
            id: id,
            name: state.product.name
          }
        }
      });
    }
    if (target.name == 'input-name') {
      this.setState((state) => {
        return {
          product: {
            id: state.product.id,
            name: target.value
          }
        }
      });
    }
  }

  render() {
    return (
      <Grid container
            direction="column"
            justify="space-around"
            alignItems="center"
            spacing={2}>
        <Grid item xs>
          <TextField  required
                      id="input-id"
                      name="input-id"
                      label="Identity"
                      variant="outlined"
                      className={this.props.classes.textField}
                      value={this.state.product.id}
                      onChange={this.handleTextFieldChange} />
        </Grid>
        <Grid item xs>
          <TextField required
                      id="input-name"
                      name="input-name"
                      label="Name"
                      variant="outlined"
                      className={this.props.classes.textField}
                      value={this.state.product.name}
                      onChange={this.handleTextFieldChange} />
        </Grid>
        <Grid item xs>
          <Button variant="contained"
                  color="primary"
                  className={this.props.classes.button}
                  onClick={this.onButtonClick}>
            Send
          </Button>
        </Grid>
      </Grid>
    )
  }
}

export default withStyles(styles)(ProductInsert);
