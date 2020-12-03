import React from 'react';
import ProductAPI from '../api/Product'
import Grid from '@material-ui/core/Grid';
import ProductInsert from '../unit/ProductInsert';
import ProductTable from '../unit/ProductTable';
import ProductUpdate from '../unit/ProductUpdate';

class MainPage extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      products: []
    }

    this.onProductUpdate = this.onProductUpdate.bind(this);
    this.handleProductUpdate = this.handleProductUpdate.bind(this);
    this.onProductInsert = this.onProductInsert.bind(this);
    this.handleProductInsert = this.handleProductInsert.bind(this);
  }

  onProductUpdate() {
    ProductAPI.fetchProducts(this.props.serverURL, this.handleProductUpdate);
  }

  handleProductUpdate(response) {
    this.setState({
      products: response.data
    });
  }

  onProductInsert(data) {
    ProductAPI.storeProduct(this.props.serverURL, data, this.handleProductInsert);
  }

  handleProductInsert(response) {
    console.log(response.status);
  }

  render() {
    return (
      <Grid container
            direction="column"
            justify="center"
            alignItems="center"
            spacing={10}>
        <Grid item xs>
          <ProductInsert onProductInsert={this.onProductInsert}/>
        </Grid>
        <Grid item xs>
          <ProductUpdate onProductUpdate={this.onProductUpdate}/>
        </Grid>
        <Grid item xs>
          <ProductTable data={this.state.products} />
        </Grid>
      </Grid>
    )
  }
}

export default MainPage;
