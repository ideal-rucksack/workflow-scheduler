create database if not exists workflow_scheduler;

use workflow_scheduler;

create table if not exists workflow (
    id int primary key auto_increment comment '主键',
    name varchar(255) not null comment '工作流名称',
    worker_id varchar(128) comment '工作流运行ID',
    description text comment '工作流描述',
    status varchar(32) not null comment '工作流状态',
    enabled tinyint(1) default false comment '是否启用',
    schedule varchar(255) not null comment '调度规则: cron->表达式',
    schedule_type varchar(32) not null comment '调度类型: cron->定时任务',
    commit_at datetime not null comment '提交时间',
    create_at datetime not null comment '创建时间',
    modify_at datetime not null comment '修改时间'
) comment '工作流表';

create table if not exists job (
    id int primary key auto_increment comment '主键',
    name varchar(255) not null comment '任务名称',
    description text comment '任务描述',
    type varchar(128) comment '任务类型',
    depend_job_id bigint comment '依赖任务id',
    plugin text comment '任务插件如: mysql_go_100',
    workflow_id bigint comment '工作流id',
    status varchar(32) not null comment '任务状态',
    script text comment '任务脚本',
    create_at datetime not null comment '创建时间',
    modify_at datetime not null comment '修改时间'
) comment '任务表';

create table if not exists job_log (
    id int primary key auto_increment comment '主键',
    job_id bigint not null comment '任务id',
    status varchar(32) not null comment '任务状态',
    log text comment '任务日志',
    create_at datetime not null comment '创建时间',
    modify_at datetime not null comment '修改时间'
) comment '任务日志表';

insert into workflow (name, description, status, enabled, schedule, schedule_type, commit_at, create_at, modify_at) values ('测试CRON', '测试一下CRON', 'waiting', 1, '*/2 * * * * *', 'cron', now(), now(), now());
insert into job (id, name, description, type, depend_job_id, workflow_id, status, script, create_at, modify_at) values (null, '代码生成', '生成代码', 'code_generation', null, null, 'waiting', null, now(), now());
