import React from 'react';
import Moment from 'react-moment';
import _ from 'lodash';
import styles from '../../styles/Views/admin/LinkList.scss';
import {PageContent} from '../../components/PageContent';
import {LinkService} from '../../services/LinkService';
import {BackButton} from '../../components/BackButton';


class EditableText extends React.Component {
  static propTypes = {
    editMode: React.PropTypes.bool,
    content: React.PropTypes.string,
    className: React.PropTypes.string,
    children: React.PropTypes.node,
    onChange: React.PropTypes.func
  };

  componentWillMount() {
    this.setState({content: this.props.content});
  }

  onChange(e) {
    this.setState({content: e.target.value});
    if (this.props.onChange) this.props.onChange(e);
  }

  render() {
    const {className, children, editMode}=this.props;
    const {content}=this.state;
    if (!editMode) return <p className={className}>{children}{content}</p>;
    return <p className={className}>{children}<input type="text" value={content} onChange={this.onChange.bind(this)}/>
    </p>;

  }
}

class LinkItem extends React.Component {
  state = {
    editMode: false
  };
  link = {
    url: '',
    description: '',
    image: ''
  };
  static propTypes = {
    link: React.PropTypes.object,
    endEditing: React.PropTypes.func
  };

  syncData() {
    const {link}=this.props;
    this.link = {
      url: link.url,
      description: link.description,
      image: link.image,
      name: link.name
    };
    if (link.id === 'new') this.setState({editMode: true});
  }

  componentWillMount() {
    this.syncData()
  }

  componentWillReceiveProps() {
    this.syncData()
  }

  goEdit() {
    this.setState({editMode: true})
  }

  changeField(fieldName) {
    return (e) => {
      this.link[fieldName] = e.target.value;
    }
  }

  async save() {
    await LinkService.save(Object.assign(this.props.link, this.link));
    this.setState({editMode: false});
    this.props.endEditing();
  }

  delLink(linkId) {
    return async() => {
      const ok = confirm('是否删除该链接');
      if (!ok) return;
      if (linkId !== 'new') await LinkService.adminDelete(linkId, true);
      this.props.endEditing();
    }
  }

  makePrivate(linkId) {
    return async() => {
      await LinkService.adminDelete(linkId);
      this.props.endEditing();
    }
  }

  makePublic(linkId) {
    return async() => {
      await LinkService.adminRecover(linkId);
      this.props.endEditing();
    }
  }


  render() {
    const link = this.props.link;
    const {name, url, image, description} =this.link;
    const {editMode}=this.state;
    return <div>
      <article className={styles.listItem}>
        <div style={styles.mainLeft}>
          <EditableText className={styles.name} content={name} editMode={editMode} onChange={this.changeField('name')}/>
          <EditableText className={styles.meta} content={url} editMode={editMode}
                        onChange={this.changeField('url')}>URL：</EditableText>
          <EditableText className={styles.meta} content={description} editMode={editMode}
                        onChange={this.changeField('description')}>描述：</EditableText>
          {!!link.updateAt &&
          <p className={styles.meta}>时间：<Moment format="YYYY年MM月DD日 HH:mm" date={link.updateAt}/></p>}
          <EditableText className={styles.meta} content={image} editMode={editMode}
                        onChange={this.changeField('image')}>图像：</EditableText>
        </div>
        <div>
          {!this.state.editMode && <button onClick={() => this.goEdit()}>编辑</button>}
          {!this.state.editMode && link.isPublic &&
          <button onClick={this.makePrivate(link.id)} style={{color: 'red'}}>设为私人</button>}
          {!this.state.editMode && !link.isPublic &&
          <button onClick={this.makePublic(link.id)} style={{color: 'green'}}>设为公开</button>}
          {this.state.editMode && <button onClick={() => this.save()}>保存</button>}
          {this.state.editMode && <button onClick={this.delLink(link.id)} style={{color: 'red'}}>彻底删除</button>}
        </div>
      </article>
    </div>;
  }
}

export class AdminLinkList extends React.Component {
  state = {
    links: []
  };

  async load() {
    const res = await LinkService.adminList();
    this.setState({links: res});
  }

  async componentWillMount() {
    document.title = '链接列表-西道の狗窝';
    await this.load();
  }

  addNew() {
    if (this.state.links[0] && this.state.links[0].id === 'new') return;
    const newLink = {
      id: 'new',
      title: '',
      url: '',
      description: '',
      image: '',
      name: ''
    };
    this.setState({links: _.concat([newLink], this.state.links)});
  }

  render() {
    return (
      <PageContent>
        <div className={styles.listWrapper}>

          <div className={styles.back}>
            <BackButton to="/"/>
          </div>
          <div className={styles.createNew} onClick={this.addNew.bind(this)}>+ 新建</div>
          {this.state.links.map((link) => <LinkItem link={link} key={link.id} endEditing={() => this.load()}/>)}

        </div>
      </PageContent>
    );
  }
}
