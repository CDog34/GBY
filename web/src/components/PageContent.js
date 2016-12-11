import React from 'react';

export class PageContent extends React.Component {
  static propTypes = {
    children: React.PropTypes.node
  };

  render() {
    return <div style={{
      height: '100vh',
      overflowY: 'auto',
      boxSizing: 'border-box'
    }}>
      {this.props.children}
    </div>
  }
}
