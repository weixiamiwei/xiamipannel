# XiamiPannel

个人浏览器导航面板（自用魔改版）。基于 [AnGe-Panel](https://github.com/liandu2024/AnGe-Panel) 二次修改，核心改动：**自建使用无需强制密码，打开即用**。

## ✨ 这版改了什么

| 改动 | 说明 |
|---|---|
| 🔓 免登录模式 | 设置环境变量 `ANGE_FREE_LOGIN=true` 后，打开面板直接进入，无需输入账号密码 |
| 🔧 开关式设计 | 不删原登录功能：不设该环境变量时，行为与 AnGe-Panel 完全一致，可正常登录管理 |
| 🚀 构建优化 | Dockerfile 启用 GOPROXY 国内代理，CI 自动构建并发布 GHCR 镜像 |
| 🐛 修复 | 统一 shell 脚本 LF 换行，修复 Windows 环境构建出的容器 entrypoint 无法执行的问题 |

### 免登录实现方式（安全说明）

- 后端登录接口忽略提交的账密，自动以**第一个启用用户（优先管理员）**身份签发 token
- 所有受保护接口自动放行为管理员身份，面板内编辑、管理功能均可正常使用
- ⚠️ **仅建议在可信局域网内使用**（家庭 NAS 场景）。免登录模式下任何能访问该地址的人都有管理员权限，请勿直接暴露到公网
- 如需恢复密码保护：去掉 `ANGE_FREE_LOGIN` 环境变量重启容器即可

## 🐳 部署

```yaml
services:
  xiamipannel:
    image: ghcr.io/weixiamiwei/xiamipannel:dev
    container_name: xiamipannel
    environment:
      - ANGE_FREE_LOGIN=true   # 免登录开关，去掉即恢复密码登录
    volumes:
      - ./data:/data
    ports:
      - "3003:3002"            # 容器内部端口取决于 data/conf/conf.ini 的 http_port
    restart: unless-stopped
```

```bash
docker compose up -d
```

数据持久化在 `./data`（conf、database、uploads、runtime），升级镜像时保留该目录即可无损升级。

支持环境：`linux/amd64`、`linux/arm64`。

## 📦 从 AnGe-Panel / Sun-Panel 迁移

直接把原面板的数据目录拷贝为 `./data` 即可，配置、收藏、图标、壁纸全部保留（本版即从 AnGe-Panel 数据原样迁移而来）。

## 🙏 致谢与版权

本项目为二次（乃至三次）开发的个人自用版本，感谢上游项目：

- [Sun-Panel](https://github.com/hslr-s/sun-panel) —— 原始项目，作者 **红烧猎人（hslr-s）**，MIT License
- [AnGe-Panel](https://github.com/liandu2024/AnGe-Panel) —— 基于 Sun-Panel v1.3.0 的增强版，作者 **liandu2024**

## 📄 License

[MIT](./LICENSE)（继承自 Sun-Panel 原始许可证，版权声明保留）。本仓库的修改部分同样以 MIT 发布。

> Sun-Panel 优秀得多，如果需要正式的密码保护与多用户支持，建议直接使用上游项目。
