

import java.util.ArrayList;

/**
 * @author linhaizeng
 * @Description TODO
 * @date 2021/10/28-11:44 下午
 */
/***
 *
 * @author linhaizeng
 * @date 2021-11-09
 *  ArrayList<String> array = new ArrayList<String>();
 *         String mail = "[0-9a-zA-Z]+@[0-9a-zA-Z]+\\.[0-9a-zA-Z]+";
 *         String pwd = "(?=.*[0-9])(?=.*[a-z])(?=.*[A-Z])[0-9a-zA-Z]{6,}";
 *         for(String[] newUser : newUsers){
 *             if(array.contains(newUser[0])){
 *                 array.add("用户名已经被注册");
 *             }else if(!newUser[1].matches(pwd)){
 *                 array.add("密码长度或格式错误");
 *             }else if(!newUser[2].matches(mail)){
 *                 array.add("邮箱格式不合法或已被注册");
 *             }else{
 *                 array.add(newUser[0]);
 *             }
 *         }
 *         return array;
 **/
public class test2 {
    public static void main(String[] args) {
        ArrayList<String> arr = new ArrayList<>();

        String test = "lindhsdsds";
        String pwd = "(?=.*[0-9])(?=.*[a-z])(?=.*[A-Z])[0-9a-zA-Z]{6,}";
        String mail = "[0-9a-zA-Z]+@[0-9a-zA-Z]+\\.[0-9a-zA-Z]+";
        if (test.matches(pwd)){
            System.out.println("1");
        } else {
            System.out.println("0");
        }
    }
}
