# coding=utf-8
"""
@author: Chen Chiheng
@date: 2023/3/16 0016 10:22
@description:
"""
import umysql

"""
数据库连接池参数：
最小连接数（池启动时，池中的连接数量）。
池中允许空闲最大连接数（当所有池中所有连接都未被使用，并且都可用时，连接的数量）。
允许被共享的连接数（当池中所有连接都已经被使用，若可以共享已经被使用的连接，此参数表示允许多少个连接被共享）。
最大连接数量。
在池中连接数不够用是是否阻塞。
单个连接允许被使用的最大次数，到达使用次数后的策略。
何时去检查已经被建立连接的状态：
    从不检查
    从池中获取时检查
    创建游标时
    执行语句时
    以上情景全部进行检查

具体接口及其实现逻辑：
    连接创建
    连接回收：
    关闭连接池：
        释放所有的连接，若此时有连接正在被使用
    封装DB API
"""

"""
并发下事务的执行：

"""


# 封装数据库连接。
class DBConn(object):

    def __init__(self, raw, maxusage):
        self.raw = raw
        self.maxusage = maxusage

    def query(self, sql, *args):
        return self.raw.query(sql, args)


# 实现数据库连接池。
class Pool(object):
    def __init__(self, host, port, user, password, db, mincached, maxcached, maxshared, maxusage, maxconnections,
                 blocking):
        self.host = host
        self.port = port
        self.user = user
        self.password = password
        self.db = db

        self._idle_conns = []  # 空闲连接。
        self.blocking = blocking  # 池中连接数不够时，是阻塞，还是抛错。
        self.maxusage = maxusage  # 若使用此参数，需要封装数据库连接已记录其被使用次数。
        self.mincached = mincached if mincached else 0  # 连接池被创建时，池中连接数量。
        self.maxcached = maxcached if maxcached else 0  # 连接池最大连接数量。
        self.maxconnections = maxconnections if maxconnections else 0  # 允许创建的最大连接数。
        # 协程共享连接没有副作用。
        self.maxshared = maxshared
        self._connections = 0  # 记录已经创建的连接数量。
        # maxcached参数调整。
        if self.maxcached and self.maxcached < self.mincached:
            self.maxcached = mincached
        # maxconnections参数调整。
        if self.maxconnections:
            if self.maxconnections < self.maxcached:
                self.maxconnections = self.maxcached
            if self.maxconnections < self.maxshared:
                self.maxconnections = self.maxshared
        # 初始化连接池。
        for i in range(self.mincached):
            self._idle_conns.append(self._create_conn())

    def _put(self, conn):
        """回收连接。"""
        # 在多线程环境中，此时应该获取互斥锁。
        try:
            self._idle_conns.append(conn)
            pass  # 若此时有等待获取连接的实体，在将连接放入连接池后应该唤醒等待获取连接的实体。
        finally:
            pass  # 释放锁

    def _create_conn(self):
        """建立新连接。"""
        raw = umysql.Connection()
        raw.connect(self.host, self.port, self.user, self.password, self.db)
        conn = DBConn(raw=raw, maxusage=self.maxusage)
        return conn

    def execute(self, sql, *args):
        try:
            conn = self._get_conn()
            results = conn.query(sql, *args)
        except Exception as err:
            raise err
        finally:
            pass
        return results

    def _get_conn(self):
        """获取连接。"""
        try:
            conn = self._idle_conns.pop(0)
            return conn
        except Exception as err:
            pass


if __name__ == '__main__':
    pool = Pool(
        host='172.16.16.59',
        port=3306,
        user='root',
        password='zhengtu#123.com',
        db='atm_1_cch',
        mincached=10,
        maxcached=20,
        maxusage=5,
        maxshared=10,
        maxconnections=20,
        blocking=True,
    )
    import gevent

    spawn = gevent.spawn(
        pool.execute, "select count(*) from roles"
    )
    gevent.joinall([spawn])
    print spawn.get()
