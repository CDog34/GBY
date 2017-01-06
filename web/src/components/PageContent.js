import React from 'react';

export class PageContent extends React.Component {
  static propTypes = {
    children: React.PropTypes.node
  };

  render() {
    return <div style={{
      flex: 1,
      overflowY: 'auto',
      boxSizing: 'border-box',
      backgroundColor: 'rgb(255, 242, 226)',
      backgroundImage: `url(${require('../images/bg.png')})`
    }}>
      {this.props.children}
    </div>
  }
}
