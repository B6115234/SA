import React, { FC, useState, useEffect } from 'react';
import Button from '@material-ui/core/Button';
import Swal from 'sweetalert2'; // alert
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme} from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save';
import TextField from '@material-ui/core/TextField';
import MenuItem from '@material-ui/core/MenuItem';
import ComponanceTable from '../Table';
import { DefaultApi } from '../../api/apis/';
import { EntCourse }  from '../../api/models/EntCourse';
import { EntSubjectType } from '../../api/models/EntSubjectType';
import { EntSubject } from '../../api/models/EntSubject';
import { Link as RouterLink } from 'react-router-dom';

const useStyles = makeStyles((theme: Theme) =>
 createStyles({
   root: {
     '& .MuiTextField-root': {
      marginTop: theme.spacing(3),
      margin: theme.spacing(3),
      width: '35ch',
     },
     display: 'flex',
     flexWrap: 'wrap',
     justifyContent: 'center',
   },
   margin: {
     margin: theme.spacing(3),
   },
   withoutLabel: {
     marginTop: theme.spacing(1),
   },
   textField: {
     marginTop: theme.spacing(1),
     margin: theme.spacing(3),
     width: '35ch',
   },
   Account: {
    marginTop: theme.spacing(1),
    marginLeft: theme.spacing(5),
  },
  textdepart: {
    margin: theme.spacing(1),
    marginTop: theme.spacing(6),
    marginRight: theme.spacing(0), 
    align: "Left",
  },
  textcourse: {
    margin: theme.spacing(1),
    marginTop: theme.spacing(6),
    marginRight: theme.spacing(2), 
    align: "Left",
  },
  textsub: {
    margin: theme.spacing(1),
    marginTop: theme.spacing(6),
    marginRight: theme.spacing(2), 
    align: "Left",
  },
  texttype: {
    margin: theme.spacing(1),
    marginTop: theme.spacing(6),
    marginRight: theme.spacing(0), 
    align: "Left",
  },
  button: {
    margin: theme.spacing(1),
    marginTop: theme.spacing(6),
    marginRight: theme.spacing(0), 
    align: "center",
  },
 }),
);

interface course_item {
      Courses: number;
      Subjects: number;
      SubjectTypes: number;
      // create_by: number;
}

const WelcomePage: FC<{}> = () => {
    const classes = useStyles();
    const api = new DefaultApi();
              
    const [courseitems, setCoursesItem] = useState<Partial<course_item>>({});

    const [courses, setCourses] = useState<EntCourse[]>([]);
    const [subjects, setSub] = useState<EntSubject[]>([]);
    const [types, setType] = useState<EntSubjectType[]>([]);
    const [loading, setLoading] =  useState(true);

    // alert setting
    const Toast = Swal.mixin({
      toast: true,
      position: 'top-end',
      showConfirmButton: false,
      timer: 3000,
      timerProgressBar: true,
      didOpen: toast => {
        toast.addEventListener('mouseenter', Swal.stopTimer);
        toast.addEventListener('mouseleave', Swal.resumeTimer);
      },
    });
  
    const getCourse = async () => {
      const res = await api.listCourse({ limit: 10, offset: 0 });
      setCourses(res);
    };
    
    const getSubject = async () => {
        const res = await api.listSubject({ limit: 10, offset: 0 });
        setSub(res);
    };

    const getType = async () => {
        const res = await api.listSubjectType({ limit: 10, offset: 0 });
        setType(res);
    };
  
    useEffect(() => {
        getCourse();
        getSubject();
        getType();
    }, [loading]);

    // set data to object course_items
    const handleChange = (
      event: React.ChangeEvent<{ name?: string; value: unknown }>,
    ) => {
      const name = event.target.name as keyof typeof WelcomePage;
      const { value } = event.target;
      setCoursesItem({ ...courseitems, [name]: value });
      console.log(courseitems);
    };

    // clear input form
    function clear() {
      setCoursesItem({});
    }

    // function save data
    function save() {
      const apiUrl = 'http://localhost:8080/api/v1/CourseItems';
      const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(courseitems),
      };

      console.log(courseitems); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

      fetch(apiUrl, requestOptions)
        .then(response => response.json())
        .then(data => {
          console.log(data);
          if (data.status === true) {
            setLoading(true);
            clear();
            Toast.fire({
              icon: 'success',
              title: 'บันทึกข้อมูลสำเร็จ',
            });
          } else {
            Toast.fire({
              icon: 'error',
              title: 'บันทึกข้อมูลไม่สำเร็จ',
            });
          }
        });
    }
    
  return (
    <Page theme={pageTheme.home}>
      <Header
        title={`ระบบลงทะเบียนเรียน`}
        subtitle="ระบบสร้างหลักสูตร"
      ><div style={{ marginLeft: 10 }}> <Button
          style={{ width: 125, height: 40, marginLeft: 130 }}
          variant="contained"
          color="primary"
          component={RouterLink}
          to="/"
      >
      ออกจากระบบ
       </Button>
      </div></Header>
      <Content>
        <ComponanceTable></ComponanceTable>

          <div>
          <form className={classes.root} noValidate autoComplete="off">
          <text className={classes.textcourse}>หลักสูตร</text >
              <TextField
                  id="course"
                  select
                  name="Courses"
                  label = "Select Course"
                  value = {courseitems.Courses  || ''}
                  variant="outlined"
                  onChange={handleChange}
              >
                 {courses.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.courseName}
                      </MenuItem>
                    );
                  })}
              </TextField>
          </form>
          </div>

          <div>
          <form className={classes.root} noValidate autoComplete="off">
          <text className={classes.textsub}>รหัสวิชา</text >
            <TextField
                  id="subject"
                  select
                  name="Subjects"
                  label = "Select Subject"
                  value = {courseitems.Subjects || ''}
                  variant="outlined"
                  onChange={handleChange}
            >
                  {subjects.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.subjectName}
                      </MenuItem>
                    );
                  })}
              </TextField>
          </form>
          </div>

          <div>
          <form className={classes.root} noValidate autoComplete="off">
          <text className={classes.texttype}>ประเภทวิชา</text >
              <TextField
                  id="type"
                  select
                  name="SubjectTypes"
                  label = "Select type"
                  value = {courseitems.SubjectTypes || ''}
                  variant="outlined"
                  onChange={handleChange}
              >
                  {types.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.typeName}
                      </MenuItem>
                    );
                  })}
              </TextField>
          </form>
          </div>

            <div className={classes.root}>
            <Button  variant="contained" size="large" startIcon = {<SaveIcon />}  onClick={save}> SAVE </Button>   
            </div>     
      </Content>
    </Page>
  );
  };
  
export default WelcomePage;
