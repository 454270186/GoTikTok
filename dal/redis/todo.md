- 当前服务节点尝试获取redis锁时，如果上锁失败，等到消息队列的解锁通知
- 在解锁时，向消息队列中发布通知，告知其他服务节点锁可用

- 使用redis一主三从，3️哨兵作为分布式锁