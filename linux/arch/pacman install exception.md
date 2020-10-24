### description
```
error: failed retrieving file 'bashtop-0.9.24-1-any.pkg.tar.zst' from mirrors.neusoft.edu.cn : The requested URL returned error: 404
error: failed retrieving file 'bashtop-0.9.24-1-any.pkg.tar.zst' from mirrors.tuna.tsinghua.edu.cn : The requested URL returned error: 404
error: failed retrieving file 'bashtop-0.9.24-1-any.pkg.tar.zst' from mirrors.celianvdb.fr : The requested URL returned error: 404
```

### resolution
force to update software source
```bash
sudo pacman -Syy
```