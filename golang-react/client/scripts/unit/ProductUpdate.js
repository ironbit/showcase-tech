import React from 'react';
import Grid from '@material-ui/core/Grid';
import Button from '@material-ui/core/Button';

class ProductUpdate extends React.Component {
  constructor(props) {
    super(props);

    this.onButtonClick = this.onButtonClick.bind(this);
  }

  onButtonClick() {
    this.props.onProductUpdate()
  }

  render() {
    return (
      <Grid item>
        <Button variant="contained"
                color="primary"
                onClick={this.onButtonClick}>
          Update Products
        </Button>
      </Grid>
    )
  }
}

export default ProductUpdate;
