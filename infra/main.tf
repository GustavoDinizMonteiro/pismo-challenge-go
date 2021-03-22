provider "aws" {
  region = "us-east-1"
}

resource "aws_lb" "monolitic_load_balance" {
  name               = "monolitic-load-nalance"
  internal           = false
  load_balancer_type = "application"
  security_groups    = [aws_security_group.lb_sg.id]
  subnets            = [aws_subnet.main_a.id, aws_subnet.main_b.id]

  enable_deletion_protection = false

  tags = {
    Environment = "dev"
  }
}

resource "aws_lb_listener" "monolitic_listiner" {
  load_balancer_arn = aws_lb.monolitic_load_balance.arn
  port              = "80"
  protocol          = "HTTP"

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.monolitic_target_group.arn
  }
}

resource "aws_lb_target_group" "monolitic_target_group" {
  name     = "monolitic-target-group"
  port     = 80
  protocol = "HTTP"
  vpc_id   = aws_vpc.main.id
}

resource "aws_vpc" "main" {
  cidr_block = "10.0.0.0/16"
}

resource "aws_internet_gateway" "main_ig" {
  vpc_id = aws_vpc.main.id

  tags = {
    Name = "main"
  }
}

resource "aws_subnet" "main_a" {
  vpc_id     = aws_vpc.main.id
  availability_zone = "us-east-1a"
  cidr_block = "10.0.0.0/24"

  tags = {
    Name = "Main"
  }
}

resource "aws_subnet" "main_b" {
  vpc_id     = aws_vpc.main.id
  availability_zone = "us-east-1b"
  cidr_block = "10.0.2.0/24"

  tags = {
    Name = "Main"
  }
}

resource "aws_alb_target_group_attachment" "monolitic_attachemnt_1" {
  target_group_arn = aws_lb_target_group.monolitic_target_group.arn
  target_id        = aws_instance.instance_1.id
  port             = 80
}

resource "aws_alb_target_group_attachment" "monolitic_attachemnt_2" {
  target_group_arn = aws_lb_target_group.monolitic_target_group.arn
  target_id        = aws_instance.instance_2.id
  port             = 80
}

resource "aws_key_pair" "terraform_user_monolitic" {
  key_name   = "terraform-user-monolitic"
  public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDcHPGq/xOE/5ITfWXSzWbrZOPEKrZ6rTupihR7w5JKfnYSa2EfVnfkn3/U4q/Lji+mvhGxDc4cCMhW1Yaj01vQKyABKBr2INgmxA97L/UAjr5PAkOOJFzR6I6fEBB4sVNPuza+wCvZBChsmZXyJQPLbDBkIj1TQmfov8z/XRBUp7oMWui2/s3AKdPxkIBs7DVQwKp8fdOOIOZRivcaZ6sOFbV8vFV8C26HZjaUog622+uE+JtmR+MihpHH7Otet55urKgtJK9r9+PDZD5mf0WRZiFPnje5Su62OcrHNPZSPNNlzPCZdc0hEnyPIrTX/bNeWn1Oj+aEevAtniTnfxlaoKSaRMvDexH7IDPlOqrQenAzyT6PnDh5ef69sg6ZdIiB0v9hwD7svfngnwU/npsMraa/IEjwpYXoWORppOI9N0SNJCYPeMbV5e1Wu2AO9ThhnJdd/ZR5B9UBiMknqVdVKuZAYzn5XoJ0eQ9NUphlYyQteB/d830T27uOuK3Skmoj7jivTnh4QSEu+gSMd+uu2rXVGO6PjKPVvQ2g4DLkt6JUydteKC8k7KVEA2G/GORUrsLno55Bpq4AflDqdwB2DhI3NNUskB7APCvJdEwxnkbrdXqHw6Flaiz9Oi1WlrDqYKBWSgo0a8mHsu2MBYUb4bT5T+Bwu6HxohVCRYWiIw== dinizmonteirogustavo@gmail.com"
}

resource "aws_security_group" "monolitic_instance_sg" {
  vpc_id = aws_vpc.main.id
  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group_rule" "monolitic_instance_sg_rule_http_ingress" {
  type              = "ingress"
  from_port         = 80
  to_port           = 80
  protocol          = "tcp"
  cidr_blocks       = ["0.0.0.0/0"]
  security_group_id = aws_security_group.monolitic_instance_sg.id
}

resource "aws_security_group" "lb_sg" {
  vpc_id = aws_vpc.main.id
  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_network_interface" "net_int" {
  subnet_id   = aws_subnet.main_a.id

  tags = {
    Name = "primary_network_interface"
  }
}

resource "aws_network_interface" "net_int_b" {
  subnet_id   = aws_subnet.main_b.id

  tags = {
    Name = "primary_network_interface"
  }
}

resource "aws_instance" "instance_1" {
  ami                    = "ami-042e8287309f5df03"
  instance_type          = "t2.micro"
  key_name               = aws_key_pair.terraform_user_monolitic.id
  tags = {
    Name = "-"
  }

  network_interface {
    network_interface_id = aws_network_interface.net_int.id
    device_index         = 0
  }
}

resource "aws_instance" "instance_2" {
  ami                    = "ami-042e8287309f5df03"
  instance_type          = "t2.micro"
  key_name               = aws_key_pair.terraform_user_monolitic.id
  tags = {
    Name = "-"
  }

  network_interface {
    network_interface_id = aws_network_interface.net_int_b.id
    device_index         = 0
  }
}