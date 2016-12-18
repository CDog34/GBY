import React from 'react';

export class PageContent extends React.Component {
  static propTypes = {
    children: React.PropTypes.node
  };

  render() {
    return <div style={{
      flex: 1,
      overflowY: 'auto',
      boxSizing: 'border-box'
    }}>
      {this.props.children}
    </div>
  }
}
